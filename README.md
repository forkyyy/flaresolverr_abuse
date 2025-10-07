# FlareSolverr — Research & Defensive Notes

> NOTE: This repository has been reoriented for defensive research, analysis, and security guidance.
> Offensive scripts and abuse tools have been kept for helping mitigate the vulnerabilities, if not, already patched, thankfully!

## Purpose
This repository contains tools used by cyber actors to abuse **FlareSolverr** instances. The focus is on **detection**, **mitigation**, and secure deployment practices to reduce the risk of abuse.

## Recommended hardening (summary)
- Enable API key authentication and rotate keys regularly.
- Restrict access via IP allowlists and reverse proxy ACLs.
- Apply rate limits per IP/session and limit concurrent browser sessions.
- Monitor logs, session counts and abnormal patterns; alert on anomalies.
- Keep FlareSolverr and components up to date.

## Responsible behaviour
- Use the materials only on systems you own or where you have explicit authorization.
- Do not use this repo to scan or attack third-party services.
- If you discover an exposed instance, notify the owner or the upstream project through a responsible disclosure channel.

## Contact
For coordination, responsible disclosure or collaboration: forkcontato@gmail.com
