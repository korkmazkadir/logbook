

const API_ROOT_URL = "http://localhost:9000/api"
const BOOK_ID = "test_1"

const LogEditorController = function () {
    this.logContentPlaceholder = "Write something to start loging..."
    this.logContent = ""
    this.logButtonDisabled = function () {
        return !(this.logContent.length > 0)
    }
    this.appendLog = function () {

        const reqBody = { content: this.logContent };
        console.log(JSON.stringify(reqBody))
        request = new Request(`${API_ROOT_URL}/${BOOK_ID}/logs`, {
            method: "POST",
            body: JSON.stringify(reqBody),
        });

        fetch(request)
            .then((response) => response.json())
            .then((json) => window.dispatchEvent(new CustomEvent("new-log", { detail: json })))
            .then(()=>this.logContent ="")
    }
}

const LogViewController = function () {
    this.placeholder = "Loading logs..."
    this.logs = []
    this.reversedLogs = function(){
        return this.logs.reverse()
    }
    this.getLogs = function () {

        console.log("loading logs...")
        fetch(`${API_ROOT_URL}/${BOOK_ID}/logs`)
            .then((response) => response.json())
            .then((json) => this.logs = json);
    }
}
