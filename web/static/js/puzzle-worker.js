
importScripts("sha256.min.js");


const leadingZeroCount = function (byteArray) {

    var zeroCount = 0;
    for (let i = 0; i < byteArray.length; i++) {
        const byte = byteArray[i]

        if (byte == 0){
            zeroCount+=8
            continue
        }

        const str = byte.toString(2).padStart(8, "0")

        for (let j = 0; j < str.length; j++) {

            if (str.charAt(j) != "0") {
                return zeroCount;
            }

            zeroCount++;
        }
    }

    return zeroCount;
}


const solvePuzzle = async function (puzzle) {

    const str = puzzle.RandomBytes
    const difficulty = puzzle.Difficulty

    var startTime, endTime;

    function start() {
      startTime = performance.now();
    };
    
    function end(message) {
      endTime = performance.now();
      var timeDiff = endTime - startTime; //in ms 
      console.log(`>>[${message}]: ${timeDiff} milliseconds`);
    }


    const encoder = new TextEncoder();
    start()
    for (let i = 0; ; i++) {

        const encoded = encoder.encode(str + i);

        const h = sha256.create()
        h.update(encoded)
        const bytes = h.array();

        if (leadingZeroCount(bytes) >= difficulty) {
            puzzle.Solution = i
            postMessage(puzzle)
            console.log(`found solution: ${i}`)
            break
        }

    }

    end("solution found")

}


onmessage = (e) => {
    const puzzle = e.data[0]
    solvePuzzle(puzzle)
};

