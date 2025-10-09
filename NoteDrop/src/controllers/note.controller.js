import { PrismaClient } from "@prisma/client";
import slugify from "slugify";
import { v4 as uuidv4 } from "uuid";

const prisma = new PrismaClient();

//Create Note
export const createNote = async (req, res) => {
    const { title, content, isPublic } = req.body;
    const { userId } = req.user;

    //input validation
    if (!title || !content) {
        return res.status(400).json({ error: "Title and content are required" });
    }

    try {
        //Generate unique slug
        const rawSlug = slugify(title, { lower: true, strict: true });
        const uniqueSlug = `${rawSlug}-${uuidv4().slice(0, 6)}`;

        //create note in DB
        const note = await prisma.note.create({
            data: {
                title,
                content,
                isPublic: isPublic || false,
                slug: uniqueSlug,
                userId,
            },
        });

        res.status(201).json({
            message: "Note created",
            note,
        });
    } catch (err) {
        console.error("Create Note Error:", err);
        res.status(500).json({ error: "Server error while creating note" });
    }
};

export const getMyNotes = async (req, res) => {
    const { userId } = req.user;

    try {
        const notes = await prisma.note.findMany({
            where: { userId },
            orderBy: { createdAt: "desc" },
        });
        res.status(200).json({
            count: notes.length,
            notes,
        });
    } catch (err) {
        console.error("Fetch My Notes Error:", err);
        res.status(500).json({ error: "Could not fetch notes" });
    }
};

export const getPublicNote = async (req, res) => {
    const { slug } = req.params;

    try {
        //Find note by slug
        const note = await prisma.note.findUnique({
            where: { slug },
        });

        //If note not found or private -> deny
        if (!note || !note.isPublic) {
            return res.status(404).json({ error: "Note not found or not public" });
        }

        //Return public note
        res.status(200).json({
            note: {
                title: note.title,
                content: note.content,
                createdAt: note.createdAt,
                slug: note.slug,
            },
        });
    } catch (err) {
        console.error("Public Note Fetch Error:", err);
        res.status(500).json({ error: "Server error while fetching note" });
    }
};

export const deleteNote = async (req, res) => {
    const { id } = req.params;
    const { userId, role } = req.user;

    try {
        //find note first
        const note = await prisma.note.findUnique({
            where: { id },
        });

        if (!note) {
            return res.status(404).json({ error: "Note not found" });
        }

        // check ownership or admin
        if (note.userId !== userId && role !== "ADMIN") {
            return res
                .status(403)
                .json({ error: "Not authorized to delete this note" });
        }
        await prisma.note.delete({
            where: { id },
        });
        res.status(200).json({ message: "Not deleted successfully" });
    } catch (err) {
        console.error("Delete Note Error:", err);
        res.status(500).json({ error: "Server error while deleting note" });
    }
};
