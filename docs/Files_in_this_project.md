Every File in nhe-mirrors, Explained

The Root Level
main.go
The front door of the entire program. In Go, every program needs exactly one main function and this is where it lives. But here's the thing — it should do almost nothing. Its only job is to start the CLI. One line. The reason you keep it empty is discipline: if you put real logic in main.go, it becomes untestable and impossible to reuse. Think of it as the ignition key. It starts the engine, it is not the engine.
go.mod
This is your project's identity card and shopping list combined. It declares the name of your module (which is how Go files refer to each other), which version of Go you're targeting, and what external packages your project depends on. When you run go get to add a dependency, it writes here. You read this file often. You rarely edit it by hand.
go.sum
This is the security guard that works with go.mod. While go.mod says "I need this package," go.sum stores a cryptographic fingerprint of every dependency so Go can verify that what you downloaded today is exactly what you downloaded last month and hasn't been tampered with. You never edit this file. You never need to read it. You always commit it to version control.
Makefile
A shortcut file. Instead of remembering the exact go build command with all its flags every time, you write it once in here and then just type make build. It also documents how the project is built, which is valuable for anyone (including future you) who picks this up later. Think of it as a recipe book for common tasks: build, install, test, cross-compile for ARM, clean up.
README.md
The document you already have. It's the first thing anyone reads. It answers: what is this, why does it exist, how do I build it, how do I use it.

cmd/root.go
This is the conductor. It doesn't know how to fetch mirrors, doesn't know how to test speed, doesn't know how to detect distros. What it knows is the sequence: get a distro, fetch its mirrors, test them, sort them, print them, optionally apply them.
It also owns the CLI surface — all the flags the user can pass (--distro, --top, --workers, --apply) are defined here. When the user types a command, this file is what interprets their intentions and delegates the actual work to the right pieces.
The reason this lives in cmd/ rather than main.go is that cmd/ can hold multiple commands later. Right now you have one command. Eventually you might add nhe-mirrors list to show supported distros, or nhe-mirrors benchmark to run extended tests. Each gets its own file in cmd/.

internal/distros/distro.go
Already covered, but to reinforce: this is the contract. The interface. The rules. Nothing in this file does work — it only defines what shape work must take. Every other file in the distros/ package is an answer to the question this file asks.
The word internal in the path is not just a folder name — in Go it has a special meaning. Files inside internal/ can only be imported by code in the same project. This is intentional. You're saying: these are implementation details, not a public API. If someone else imports your module, they get the CLI binary — they don't get to reach into your internals.
internal/distros/registry.go
Already covered. The phonebook and the auto-detector.

The Distro Files
Each of these is the same shape — they all implement the same interface — but each one knows intimate details about one specific distribution. Think of them as specialists. arch.go is the Arch specialist. It knows where Arch publishes its mirror list, what format that list is in, what file to download to test speed, and exactly how a valid Arch mirror entry looks in pacman.conf. It knows nothing about Debian. It doesn't need to.
arch.go
The cleanest implementation because Arch publishes a proper JSON API for their mirrors. It gives you the URL, the country, the protocol, and even their own health score you can use to pre-filter bad mirrors before you even test them. Build this first.
debian.go
Debian also has a JSON API but the structure is different from Arch. The output format is different too — an apt sources line looks nothing like a pacman server line. Same interface, completely different internals.
ubuntu.go
Ubuntu's mirror list is the simplest to parse: plain text, one URL per line. But there's a subtlety — Ubuntu has release codenames (noble, jammy, focal) and the sources line needs to know which one you're on. So this implementation needs to be aware of the release version, which it reads either from a flag or from the running system.
alpine.go
Alpine is interesting because it's a musl-based distro — lighter than the others — and its mirror format and repository structure are different. It also has a JSON mirror list. The config file format (/etc/apk/repositories) is simpler than apt but the path structure inside the mirror is different.
fedora.go
Fedora is the odd one. It uses a metalink format by default, which is XML that points to multiple mirrors with checksums. For your purposes you bypass that and go directly to the mirrorlist endpoint, which gives you plain text. The output format — a .repo file — is INI-style, which is different from everything else.
rocky.go and almalinux.go
These are the spiritual successors to CentOS. They use the same dnf package manager as Fedora and the same .repo file format. Their mirror list endpoints are different URLs but the parsing and output logic is nearly identical to Fedora. Once you've written fedora.go, these are fast to add.
void.go
Void is unusual — it doesn't have a polished JSON API like Arch. The mirror list lives in a raw text file in their GitHub repository. So this implementation fetches a file from GitHub, parses plain text, and writes to a different config path than everyone else. It teaches you the third and final parsing pattern: "just a list of URLs in a file somewhere."
opensuse.go
openSUSE uses zypper and has its own repository structure. The mirror list is published on their website. The config output goes to /etc/zypp/repos.d/ which is a directory rather than a single file — a pattern you don't see in the other distros.

internal/tester/tester.go
This is the engine room. It knows nothing about any distro — it doesn't know what Arch is, doesn't know what apt is. All it knows is: "give me a list of URLs and a file path, I will download that file from every URL simultaneously and tell you how fast each one was."
It has two moving parts. The first is the function that tests a single mirror: make an HTTP request, record the time, drain the response body, calculate bytes per second. The second is the function that runs the first one across hundreds of mirrors at once using goroutines — but with a limiter so you don't open 300 simultaneous connections and get banned or crash.
The reason this is separate from the distro files is the single responsibility principle. Testing speed is a generic, reusable capability. It has nothing to do with what distro you're running. Keeping it separate means you could theoretically use this tester for something entirely different someday, and it means you can test it in isolation without needing any distro logic.

How They All Relate
main.go
  └── starts cmd/root.go
        ├── asks registry.go   "which distro am I on?"
        │     └── registry returns a distro (arch.go, ubuntu.go, etc.)
        ├── calls distro.FetchMirrors()
        │     └── that distro implementation fetches its own list
        ├── passes mirrors to tester.go
        │     └── tester fills in Speed and Latency on each Mirror
        ├── sorts by Speed (logic defined in distro.go)
        └── prints results / writes config using distro.FormatEntry()
Every file has exactly one reason to exist. That's the goal.