# migrant
Migrate to esa.io from your wiki tool.

## Usage

### Set up

1. Install go (> 1.13)
1. Clone this repos.
1. `go build .`

### Arguments

To use migrant, please set arguments correctly.

- `--team "TEAM_NAME_HERE"` (Required)
    - Team name in your esa.io team.
- `--key "API_KEY_HERE"` (Required)
    - Personal Access Token in your esa.io team.
- `--max-posts 75` (Optional, default 70)
    - The number of posts within window size(wrote below).
    - More info of rate limit can found in [official esa.io docs](https://docs.esa.io/posts/102) .
- `--window-size 15` (Optional, default 15)
    - Waiting to reset rate limit these minutes.
    - NOTE : migrant is NOT read `X-RateLimit-Reset` or `429 Too Many Requests` .

### Migrate from esa.io (exported by esa.io)

1. Put `backup` dir to working dirs.
1. Put exported data into `backup` .
    - ex: `backup/README.md`, `backup/SOMEDIRS/awesome_file.md`
1. `./migrant --team "{Team_Name}" --key "{Personal_Access_Token}"`
    - NOTE : Imported files put in `backup` category. migrant refer to the dirs structure.
    - NOTE : migrant can NOT read tags, comments, wip, and so on...
        - Did't read the yaml called "Front Matter"


## Work in Progress

- Compatible with esa.io (Team to Team, through API)
- Compatible with Crowi / Growi

More info of WIP, please read issues.