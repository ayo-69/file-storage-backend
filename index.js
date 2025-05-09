const express = require("express");
const app = express();

app.get("/", (req, res) => {
    res.status(200).json({
        message: "Hello world"
    })
})

const PORT = process.env.PORT | 3030;
app.listen(PORT, () => {
    console.log(`Server running on PORT : ${ PORT }`);
});
