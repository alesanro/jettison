# Changelog

All notable changes to this project are documented in this file.

**THIS LIBRARY IS STILL IN ALPHA AND THERE ARE NO GUARANTEES REGARDING API STABILITY YET**

## [v0.2.1] - 2019-09-10
- Refactor instructions for types implementing the `json.Marshaler` and `encoding.TextMarshaler` interfaces.
   - Fix encoding of `nil` instances.
   - Fix behavior for pointer and non-pointer receivers, to comply with `encoding/json`.
- Fix bug that prevents tagged fields to dominate untagged fields.
- Add support for anonymous struct pointer fields.
- Improve tests coverage of `encoder.go`.
- Add test cases for unexported non-embedded struct fields.

## [v0.2.0] - 2019-09-01
- Add support for `json.Number`.
- Update `README.md` to add a Go1.12+ requirement notice.

## [v0.1.0] - 2019-08-30
Initial realease.

[v0.2.1]: https://github.com/wI2L/jettison/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/wI2L/jettison/compare/0.1.0...v0.2.0
[v0.1.0]: https://github.com/wI2L/jettison/releases/tag/0.1.0