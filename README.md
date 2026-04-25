# AutoUnix (AU) 🚀
AutoUnix is an automation tool for online courses, powered by Go and Generative AI.

## Tech Stack
* **Language:** [Go (Golang)](https://go.dev/) 1.21+
* **Browser automatisation:** [Rod](https://go-rod.github.io/) (DevTools driver).
* **AI integration:** Google Gemini API.
* **Parsing:** DOM scraping via CSS and XPath selectors.

## Project structure
* `main.go` — The entry point and core application logic..
* `internal/browser/` — Browser initialization and instance management.
* `internal/ai/` — AI interaction(Gemini).
* `internal/parser/` — Question parsing, site navigation, and interaction logic.

## 🚀 Getting started
### 1. Installation
Clone the repository to your local machine:
```bash
git clone https://github.com/Pirot1/AutoUnix.git
cd AutoUnix
```
### 2. Dependencies
Install the required Go modules:
```bash
go mod tidy
```
### 3. Configuration
The .env file will create automatically with first launch
### 4. Running the Bot 
Execute the source code:
```bash
go run ./cmd/bot/main.go
```
Or run the pre-compiled binary:
```bash
./autounix.exe
```

---

## 🏗 Project Architecture 

The project follows a modular structure to ensure clean code separation and scalability:



```text
.
├── cmd/
│   └── bot/
│       └── main.go           # Entry point (Application startup)
├── internal/
│   ├── ai/
│   │   ├── gemini.go         # API-based answer selection (multiple choice 1-4)
│   │   ├── web_gemini.go     # Web-based interaction with Gemini 
│   │   └── quiz.go           # Quiz processing and test execution
│   ├── browser/
│   │   └── init.go           # Browser initialization (Headless mode, etc.)
│   └── parser/
│       ├── authentication.go # Site login and session management
│       ├── caption.go        # Capturing subtitles and generating summaries
│       ├── lesson_finder.go  # Logic to locate specific course modules
│       ├── sidebar.go        # Progress tracking and lesson navigation
│       └── video_watcher.go  # Video playback control and monitoring
├── lessons/
│   └── Lesson_name/            # Directory for stored study notes/summaries
├── .env                      # Secret keys and credentials (Private)
├── env.example               # Template for the .env file
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
├── .gitignore                # Files excluded from version control
└── README.md                 # Project documentation (You are here)
```
## Key Features
* **Smart Summaries**: Automatically extracts lesson subtitles and uses AI to generate concise, structured study notes.

* **AI Quiz Solver**: Parses multiple-choice questions and consults Gemini to determine the most accurate answer.

* **Headless Execution**: Runs silently in the background, allowing you to use your PC while the bot works.

* **Intelligent Navigation**: Automatically detects completed lessons and navigates through the course sidebar.

## Development Note
`For cross-platform compatibility, it is recommended to keep directory names in English to avoid potential encoding issues during automated file operations.`