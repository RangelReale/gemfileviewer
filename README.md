# gemfileviewer

```shell
go install github.com/RangelReale/gemfileviewer/cmd/gemfileviewer@latest
```

```shell
gemfileviewer [-n N] <file.gem>
```

The optional parameter `-n` outputs the file numbered `N` at the end.

Sample output:

```
====== METADATA (metadata.gz) ======
--- !ruby/object:Gem::Specification
name: thor
version: !ruby/object:Gem::Version
  version: 1.2.1
platform: ruby
authors:
- Yehuda Katz
- José Valim
autorequire: 
bindir: bin
cert_chain: []
date: 2022-01-04 00:00:00.000000000 Z
dependencies:
- !ruby/object:Gem::Dependency
  name: bundler
  requirement: !ruby/object:Gem::Requirement
    requirements:
    - - ">="
      - !ruby/object:Gem::Version
        version: '1.0'
    - - "<"
      - !ruby/object:Gem::Version
        version: '3'
  type: :development
  prerelease: false
  version_requirements: !ruby/object:Gem::Requirement
    requirements:
    - - ">="
      - !ruby/object:Gem::Version
        version: '1.0'
    - - "<"
      - !ruby/object:Gem::Version
        version: '3'
description: Thor is a toolkit for building powerful command-line interfaces.
email: ruby-thor@googlegroups.com
executables:
- thor
extensions: []
extra_rdoc_files: []
files:
- ".document"
- CONTRIBUTING.md
- LICENSE.md
- README.md
- bin/thor
- lib/thor.rb
- lib/thor/actions.rb
- lib/thor/actions/create_file.rb
- lib/thor/actions/create_link.rb
- lib/thor/actions/directory.rb
- lib/thor/actions/empty_directory.rb
- lib/thor/actions/file_manipulation.rb
- lib/thor/actions/inject_into_file.rb
- lib/thor/base.rb
- lib/thor/command.rb
- lib/thor/core_ext/hash_with_indifferent_access.rb
- lib/thor/error.rb
- lib/thor/group.rb
- lib/thor/invocation.rb
- lib/thor/line_editor.rb
- lib/thor/line_editor/basic.rb
- lib/thor/line_editor/readline.rb
- lib/thor/nested_context.rb
- lib/thor/parser.rb
- lib/thor/parser/argument.rb
- lib/thor/parser/arguments.rb
- lib/thor/parser/option.rb
- lib/thor/parser/options.rb
- lib/thor/rake_compat.rb
- lib/thor/runner.rb
- lib/thor/shell.rb
- lib/thor/shell/basic.rb
- lib/thor/shell/color.rb
- lib/thor/shell/html.rb
- lib/thor/util.rb
- lib/thor/version.rb
- thor.gemspec
homepage: http://whatisthor.com/
licenses:
- MIT
metadata:
  bug_tracker_uri: https://github.com/rails/thor/issues
  changelog_uri: https://github.com/rails/thor/releases/tag/v1.2.1
  documentation_uri: http://whatisthor.com/
  source_code_uri: https://github.com/rails/thor/tree/v1.2.1
  wiki_uri: https://github.com/rails/thor/wiki
  rubygems_mfa_required: 'true'
post_install_message: 
rdoc_options: []
require_paths:
- lib
required_ruby_version: !ruby/object:Gem::Requirement
  requirements:
  - - ">="
    - !ruby/object:Gem::Version
      version: 2.0.0
required_rubygems_version: !ruby/object:Gem::Requirement
  requirements:
  - - ">="
    - !ruby/object:Gem::Version
      version: 1.3.5
requirements: []
rubygems_version: 3.2.32
signing_key: 
specification_version: 4
summary: Thor is a toolkit for building powerful command-line interfaces.
test_files: []

====== DATA (data.tar.gz) ======
[1] .document
[2] CONTRIBUTING.md
[3] LICENSE.md
[4] README.md
[5] bin/thor
[6] lib/thor.rb
[7] lib/thor/actions.rb
[8] lib/thor/actions/create_file.rb
[9] lib/thor/actions/create_link.rb
[10] lib/thor/actions/directory.rb
[11] lib/thor/actions/empty_directory.rb
[12] lib/thor/actions/file_manipulation.rb
[13] lib/thor/actions/inject_into_file.rb
[14] lib/thor/base.rb
[15] lib/thor/command.rb
[16] lib/thor/core_ext/hash_with_indifferent_access.rb
[17] lib/thor/error.rb
[18] lib/thor/group.rb
[19] lib/thor/invocation.rb
[20] lib/thor/line_editor.rb
[21] lib/thor/line_editor/basic.rb
[22] lib/thor/line_editor/readline.rb
[23] lib/thor/nested_context.rb
[24] lib/thor/parser.rb
[25] lib/thor/parser/argument.rb
[26] lib/thor/parser/arguments.rb
[27] lib/thor/parser/option.rb
[28] lib/thor/parser/options.rb
[29] lib/thor/rake_compat.rb
[30] lib/thor/runner.rb
[31] lib/thor/shell.rb
[32] lib/thor/shell/basic.rb
[33] lib/thor/shell/color.rb
[34] lib/thor/shell/html.rb
[35] lib/thor/util.rb
[36] lib/thor/version.rb
[37] thor.gemspec

====== GEMSPEC ======
# coding: utf-8
lib = File.expand_path("../lib/", __FILE__)
$LOAD_PATH.unshift lib unless $LOAD_PATH.include?(lib)
require "thor/version"

Gem::Specification.new do |spec|
  spec.add_development_dependency "bundler", ">= 1.0", "< 3"
  spec.authors = ["Yehuda Katz", "José Valim"]
  spec.description = "Thor is a toolkit for building powerful command-line interfaces."
  spec.email = "ruby-thor@googlegroups.com"
  spec.executables = %w(thor)
  spec.files = %w(.document thor.gemspec) + Dir["*.md", "bin/*", "lib/**/*.rb"]
  spec.homepage = "http://whatisthor.com/"
  spec.licenses = %w(MIT)
  spec.name = "thor"
  spec.metadata = {
    "bug_tracker_uri" => "https://github.com/rails/thor/issues",
    "changelog_uri" => "https://github.com/rails/thor/releases/tag/v#{Thor::VERSION}",
    "documentation_uri" => "http://whatisthor.com/",
    "source_code_uri" => "https://github.com/rails/thor/tree/v#{Thor::VERSION}",
    "wiki_uri" => "https://github.com/rails/thor/wiki",
    "rubygems_mfa_required" => "true",
  }
  spec.require_paths = %w(lib)
  spec.required_ruby_version = ">= 2.0.0"
  spec.required_rubygems_version = ">= 1.3.5"
  spec.summary = spec.description
  spec.version = Thor::VERSION
end

====== FILE OUTPUT (lib/thor/version.rb) ======
class Thor
  VERSION = "1.2.1"
end
```

## Author

Rangel Reale (rangelreale@gmail.com)
