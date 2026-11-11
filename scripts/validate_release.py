import argparse
import re
import subprocess
from typing import TypedDict


class Args(TypedDict):
    tag: str


def main():
    args = get_args()

    version = args["tag"].lstrip("v")

    validate_changelog(version)
    validate_lip_toml(version)


def get_args() -> Args:
    parser = argparse.ArgumentParser()
    parser.add_argument("--tag", required=True)

    args = parser.parse_args()

    return {
        "tag": args.tag,
    }


def validate_changelog(version: str):
    try:
        subprocess.run(
            f"npx changelog --format markdownlint",
            shell=True,
            check=True,
        )
    except subprocess.CalledProcessError as e:
        print("Have you installed it by `npm i -g keep-a-changelog`?")
        raise e

    with open("CHANGELOG.md", "r", encoding="utf-8") as f:
        content = f.read()

    if not re.search(r"## \[{}\]".format(version), content):
        raise Exception("CHANGELOG.md lacks version {}".format(version))


def validate_lip_toml(version: str):
    with open("lip.toml", "r", encoding="utf-8") as f:
        content = f.read()

    match = re.search(r'^version = "(.+)"$', content, re.MULTILINE)
    if match is None:
        raise Exception("lip.toml lacks version")

    if match.group(1) != version:
        raise Exception(
            "lip.toml version {} does not match tag {}".format(match.group(1), version)
        )


if __name__ == "__main__":
    main()
