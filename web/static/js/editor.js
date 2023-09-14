
const API_ROOT_URL = `/api`
const BOOK_ID = new URLSearchParams(location.search).get("book_id")

const LogEditorController = function () {
    this.logContentPlaceholder = "Write something to start loging..."
    this.logContent = ""
    this.logButtonDisabled = function () {
        return !(this.logContent.length > 0)
    }
    this.appendLog = function () {

        const reqBody = { content: this.logContent };
        const request = new Request(`${API_ROOT_URL}/${BOOK_ID}/logs`, {
            method: "POST",
            body: JSON.stringify(reqBody),
        });

        fetch(request)
            .then((response) => response.json())
            .then((json) => window.dispatchEvent(new CustomEvent("new-log", { detail: json })))
            .then(() => this.logContent = "")
    }
}

const LogViewController = function () {
    this.logs = []
    this.reversedLogs = function () {
        return this.logs.reverse()
    }
    this.getLogs = function () {

        console.log(`Loading logs: ${API_ROOT_URL}/${BOOK_ID}/logs`)
        fetch(`${API_ROOT_URL}/${BOOK_ID}/logs`)
            .then((response) => response.json())
            .then((json) => this.logs = json ? json : []);
    }
    this.dateString = function (logTime) {
        const date = new Date(logTime)
        //return `${date.toLocaleDateString()} /  ${date.toLocaleString('en-us', {  weekday: 'long' })} / ${date.toLocaleTimeString()}`
        return `${date.toLocaleDateString()}  ${date.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}`
    }
}
