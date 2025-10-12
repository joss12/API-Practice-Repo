const sendConfirmation = require("../utils/sendEmail");
const validateForm = require("../utils/validateForm");
const path = require("path");
const fs = require("fs");

//const DATA_PATH = path.join(__dirname, "..", "data.json");
const DATA_PATH = path.join(__dirname, "..", "..", "data.json");
const submitMiddleware = async (req, res) => {
    try {
        const { name, email, message } = req.body;

        // Validate form data
        const { valid, errors } = validateForm(name, email, message);
        if (!valid) {
            return res.status(400).json({ errors });
        }

        // Create new submission
        const newSubmission = {
            id: Date.now(),
            name,
            email,
            message,
            date: new Date().toString(),
        };

        // Save to file - create file if it doesn't exist
        let submissions = [];
        if (fs.existsSync(DATA_PATH)) {
            const file = fs.readFileSync(DATA_PATH, "utf-8");
            submissions = file.trim() ? JSON.parse(file) : [];
        }
        submissions.push(newSubmission);
        fs.writeFileSync(DATA_PATH, JSON.stringify(submissions, null, 2));

        // Send email confirmation (only if not in test environment)
        if (process.env.NODE_ENV !== "test") {
            await sendConfirmation(email, name, message);
        }

        // Return success response
        return res.status(200).json({
            success: true,
            message: "Form submitted successfully",
            data: newSubmission,
        });
    } catch (error) {
        console.error("Error in submitMiddleware:", error);
        return res.status(500).json({
            success: false,
            error: "Internal server error",
            message: error.message,
        });
    }
};

const mySubmit = async (req, res) => {
    try {
        const email = req.user.email;

        if (!fs.existsSync(DATA_PATH)) return res.json([]);
        const raw = await fs.promises.readFile(DATA_PATH, "utf-8");
        const all = raw.trim() ? JSON.parse(raw) : [];

        const userSubmissions = all.filter((entry) => entry.email === email);
        res.json(userSubmissions);
    } catch (error) {
        console.error(error);
        res.status(500).json({ error: "Failed to retrieve submissions" });
    }
};

const deleteUser = (req, res) => {
    try {
        const id = Number(req.params.id); // Convert to number

        if (!fs.existsSync(DATA_PATH)) {
            return res.status(404).json({ error: "Data file not found" });
        }

        const raw = fs.readFileSync(DATA_PATH, "utf-8");
        const all = raw.trim() ? JSON.parse(raw) : [];

        const index = all.findIndex((entry) => entry.id === id);
        if (index === -1) {
            return res.status(404).json({ error: "Data not found" });
        }

        const deleted = all.splice(index, 1);
        fs.writeFileSync(DATA_PATH, JSON.stringify(all, null, 2));

        return res.json({
            success: true,
            message: `Submission ${id} deleted.`,
            deleted,
        });
    } catch (err) {
        console.error(err);
        return res.status(500).json({ error: "Server error" });
    }
};

module.exports = { submitMiddleware, mySubmit, deleteUser };
