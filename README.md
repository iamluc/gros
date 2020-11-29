gros (Github Release Object Storage)
------------------------------------

Download a file from a Github release in a private repository.

# Why?

[gh](https://github.com/cli/cli) is better (really), but...
- There is no build for FreeBSD
- Released binaries are zipped, you have to unzip them before using them

# Usage

```
export GITHUB_TOKEN=XXXXX
./gros my_org/my_private_repo release_tag file_to_download local_file_path
```
