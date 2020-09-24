const express = require('express');
const path = require('path');
const fs = require('fs');
const config = require('./config.json');
const app = express();
var fileUpload = require('express-fileupload');

function genName(ext) {
    var result = "";
    var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    var charactersLength = characters.length;
    
    for (var i = 0; i < 5; i++) 
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    
    if (fs.existsSync(path.join(__dirname, "/files/" + result + '.' + ext)))
        result = genName(ext);

    return result + '.' + ext;
}

function uploadFile(req, res) {
    if (!req.headers.authorization) {
        return res.status(403).json({ success: false, url: "", error: "No authorization sent." });
    }
    else if (req.headers.authorization != config.key) {
        return res.status(403).json({ success: false, url: "", error: "Invalid authorization." });
    }

    if (req.files) {
        if (req.files.file) {
            let ext = req.files.file.name.split('.')[1] || "";
            var name = genName(ext);

            req.files.file.mv(path.join(__dirname, "/files/" + name));

            return res.status(201).json({ success: true, url: req.hostname + '/' + name, error: "" });
        } 
        else {
            return res.status(400).json({ success: false, url: "", error: "No file named 'file' was uploaded." });
        }
    } 
    else {
        return res.status(400).json({ success: false, url: "", error: "No file uploaded." });
    }
}

function tryGetFile(req, res) {
    var file = path.join(__dirname, "/files/" + req.params.file);

    if (fs.existsSync(file)) {
        res.sendFile(file);
    }
    else {
        res.status(404).send("404. Not found!");
    }
}

app.use(fileUpload({
    safeFileNames: false,
    preserveExtension: true
}));

app.post("/upload", uploadFile);

app.get("/i/:file", tryGetFile);

app.use('/', express.static(path.join(__dirname + "/public")));

app.all('*', (req, res) => {
    res.status(404).send("404. Not found!");
});

const listener = app.listen(8080, () => {
    console.log("clyx-node >> Listening on port: " + listener.address().port);
});