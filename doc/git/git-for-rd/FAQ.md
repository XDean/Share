# FAQ


Q: When use HTTPS protocol, Git always ask to input password.
A: Run `$ git config --global credential.helper cache`, it will cache your password for 15 min. Or run `$ git config --global credential.helper store` to store your password forever. See [Git - Credential Storage](https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage)

Q: When use HTTPS protocol, Git told me SSL verify failed.
A: Run `git config --global https.sslVerify "false"` to skip SSL verify. See [Git - http sslVerify](https://git-scm.com/docs/git-config#Documentation/git-config.txt-httpsslVerify)