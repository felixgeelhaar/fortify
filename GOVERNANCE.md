# Project Governance

## Maintainership

Fortify is currently maintained by a single person: [@felixgeelhaar](https://github.com/felixgeelhaar). The project is in solo-maintainer mode; bus factor is **1**.

If you depend on Fortify in production, plan accordingly:

- Pin the minor version in your `go.mod` (e.g., `v1.2.x`, not `latest`).
- Vendor the module if your build pipeline supports it.
- Review the [SECURITY.md](SECURITY.md) disclosure process before reporting vulnerabilities.

## Decision-making

Until Fortify reaches at least three active maintainers:

- **Lazy consensus:** non-trivial design decisions are discussed in GitHub issues. If no objections are raised within 7 days, the proposal is accepted.
- **Maintainer call:** the maintainer has final authority on merges, releases, and roadmap.
- **External contributors** are welcome and credited; sustained contribution is a path to additional maintainership.

## Roadmap & priorities

The roadmap is **not** versioned in this repo. Priorities are tracked in [GitHub Issues](https://github.com/felixgeelhaar/fortify/issues) under the labels:

- `priority/high` — being worked on this milestone.
- `priority/medium` — accepted, scheduled.
- `priority/low` — accepted, no timeline.
- `idea` — under discussion, may be rejected.

Open a Discussion or Issue to propose new directions before opening a large PR.

## API stability

Fortify follows [Semantic Versioning](https://semver.org/) for the `v1.x.x` line:

- **Major (vX.0.0):** breaking API changes. Migration notes in [CHANGELOG.md](CHANGELOG.md) and `docs/MIGRATION.md`.
- **Minor (v1.X.0):** additive, backwards-compatible changes.
- **Patch (v1.x.X):** bug fixes only.

The current `v1.x` line is **iterating**. Several breaking changes have shipped recently as part of pre-stability hardening. Expect a `v1.x` release marked "stable" in the CHANGELOG when the API is frozen.

The `v2` module path is **reserved** and should not be considered active. See `docs/MIGRATION.md` for the v1↔v2 history.

## Security

Vulnerability reports go to **felix.geelhaar@gmail.com**, not GitHub Issues. See [SECURITY.md](SECURITY.md) for the full disclosure policy.

Releases include:

- CycloneDX (`sbom.cdx.json`) and SPDX (`sbom.spdx.json`) SBOMs.
- nox security scan results published as SARIF on the GitHub Security tab.
- Dependabot-managed dependency updates (gomod + GitHub Actions).
- All third-party Actions pinned to commit SHAs.

## Releases

- Cut from `main` with a signed git tag matching `v1.x.x`.
- The `release.yml` workflow runs the full test suite, generates SBOMs, and creates a GitHub Release.
- Release notes are extracted from `CHANGELOG.md` for the matching version.

## Adding maintainers

A new maintainer is added when:

1. They have made at least three substantive merged contributions (not typo fixes).
2. They have demonstrated familiarity with the architecture (reviewed at least one non-trivial PR).
3. The existing maintainer(s) agree.

Maintainer expectations:

- Triage incoming issues within ~1 week.
- Review PRs they are tagged on within ~2 weeks.
- Adhere to the security disclosure process.

## Code of Conduct

Be kind. Disagreements are expected; personal attacks, harassment, or discrimination are not. The maintainer may close discussions or remove participants who violate this norm.

## License

Fortify is MIT-licensed. See [LICENSE](LICENSE).
