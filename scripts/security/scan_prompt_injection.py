#!/usr/bin/env python3
"""Fail when repository text contains common prompt-injection instructions."""

from __future__ import annotations

import argparse
import re
import sys
from pathlib import Path


TEXT_SUFFIXES = {
    ".birl", ".go", ".html", ".json", ".md", ".txt", ".yaml", ".yml",
}
IGNORED_PARTS = {".git", "bin", "dist", "vendor"}
SELF = Path(__file__).resolve()

# Keep fragments separate so this scanner does not flag its own source.
PATTERNS = (
    ("instruction override", r"ignore\s+(?:all\s+)?previous\s+" + r"instructions?"),
    ("instruction override", r"disregard\s+(?:all\s+)?(?:prior|previous)\s+" + r"instructions?"),
    ("system prompt extraction", r"(?:reveal|print|show|leak)\s+(?:the\s+)?" + r"system\s+prompt"),
    ("role escalation", r"you\s+are\s+now\s+(?:in\s+)?" + r"(?:developer|system|admin)\s+mode"),
    ("secret extraction", r"(?:exfiltrate|reveal|print|leak)\s+(?:all\s+)?" + r"(?:secrets?|credentials?|api[_ -]?keys?)"),
)


def candidate_files(root: Path):
    for path in root.rglob("*"):
        if not path.is_file() or path.resolve() == SELF:
            continue
        if any(part in IGNORED_PARTS for part in path.parts):
            continue
        if path.suffix.lower() in TEXT_SUFFIXES:
            yield path


def scan(root: Path) -> list[str]:
    findings: list[str] = []
    compiled = [(name, re.compile(pattern, re.IGNORECASE)) for name, pattern in PATTERNS]
    for path in candidate_files(root):
        try:
            lines = path.read_text(encoding="utf-8").splitlines()
        except (OSError, UnicodeDecodeError) as error:
            findings.append(f"{path}: unable to scan text: {error}")
            continue
        for number, line in enumerate(lines, 1):
            for name, pattern in compiled:
                if pattern.search(line):
                    findings.append(f"{path}:{number}: {name}")
    return findings


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("path", nargs="?", default=".", type=Path)
    args = parser.parse_args()
    findings = scan(args.path.resolve())
    if findings:
        print("Potential prompt injection detected:", file=sys.stderr)
        print("\n".join(findings), file=sys.stderr)
        return 1
    print("Prompt-injection scan passed")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
