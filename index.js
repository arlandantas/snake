const go = new Go();
WebAssembly.instantiateStreaming(fetch("snake.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    // console.log(result);
});