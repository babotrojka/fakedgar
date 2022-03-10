    async function compile() {
        console.log(document.getElementById("input").value)
        await fetch("../calculate", {
                method: "PUT",
                body: JSON.stringify(input, document.getElementById("input").value),
            })
            .then((response => response.text()))
            .then((data) => {
                console.log(data);
            })
            .catch((error) => {
                console.error("Error:", error);
            })
    }
    document.getElementById("run").onclick = compile;