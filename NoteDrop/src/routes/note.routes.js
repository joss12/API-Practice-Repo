import express from "express";
import { protect } from "../middlewares/auth.js";
import {
    createNote,
    deleteNote,
    getMyNotes,
    getPublicNote,
} from "../controllers/note.controller.js";

const router = express.Router();

//create a new note
router.post("/", protect, createNote);

// Get own notes
router.get("/me", protect, getMyNotes);

//Get plublic note by slug
router.get("/:slug", getPublicNote);

//delete note (only owner or admin)
router.delete("/:id", protect, deleteNote);

export default router;
