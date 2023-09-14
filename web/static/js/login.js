const API_ROOT_URL = "/api"
const puzzleWorker = new Worker("static/js/puzzle-worker.js");

const LoginController = function () {
    this.message = ""
    this.url = ""
    this.spinnerVisible = false
    this.buttonVisible = true
    this.host = window.location.protocol + '//' + window.location.host
    this.createLogbook = function () {

        this.message = "Creating a new logbook may take around 2 minutes."
        this.spinnerVisible = true
        this.buttonVisible = false
        

        const createPuzzleReq = new Request(`${API_ROOT_URL}/puzzle`, {
            method: "GET",
        });

        fetch(createPuzzleReq)
            .then((response) => response.json())
            .then((puzzle) => {
                console.log(`a new puzzle retreived ${JSON.stringify(puzzle)}`)
                puzzleWorker.postMessage([puzzle])
            })


        puzzleWorker.onmessage = (e) => {
            const puzzle = e.data;
            const request = new Request(`${API_ROOT_URL}/book`, {
                method: "POST",
                body:JSON.stringify(puzzle)
            })

            fetch(request)
                .then((response) => response.json())
                .then((result) => {
                    this.url = result.url
                    this.message = "Save this URL aside to access your logbook."
                    this.spinnerVisible = false;
                });
        };

    }
}