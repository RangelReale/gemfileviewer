# gemfileviewer

```shell
go install github.com/RangelReale/gemfileviewer/cmd/gemfileviewer@latest
```

```shell
gemfileviewer <file.gem>
```

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
.document
CONTRIBUTING.md
LICENSE.md
README.md
bin/thor
lib/thor.rb
lib/thor/actions.rb
lib/thor/actions/create_file.rb
lib/thor/actions/create_link.rb
lib/thor/actions/directory.rb
lib/thor/actions/empty_directory.rb
lib/thor/actions/file_manipulation.rb
lib/thor/actions/inject_into_file.rb
lib/thor/base.rb
lib/thor/command.rb
lib/thor/core_ext/hash_with_indifferent_access.rb
lib/thor/error.rb
lib/thor/group.rb
lib/thor/invocation.rb
lib/thor/line_editor.rb
lib/thor/line_editor/basic.rb
lib/thor/line_editor/readline.rb
lib/thor/nested_context.rb
lib/thor/parser.rb
lib/thor/parser/argument.rb
lib/thor/parser/arguments.rb
lib/thor/parser/option.rb
lib/thor/parser/options.rb
lib/thor/rake_compat.rb
lib/thor/runner.rb
lib/thor/shell.rb
lib/thor/shell/basic.rb
lib/thor/shell/color.rb
lib/thor/shell/html.rb
lib/thor/util.rb
lib/thor/version.rb
thor.gemspec

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
```

## Author

Rangel Reale (rangelreale@gmail.com)
