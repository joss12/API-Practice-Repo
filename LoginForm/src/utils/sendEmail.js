const nodemailer = require("nodemailer");
require("dotenv").config();

const transporter = nodemailer.createTransport({
    host: process.env.EMAIL_HOST,
    port: process.env.EMAIL_PORT,
    secure: false,
    auth: {
        user: process.env.EMAIL_USER,
        pass: process.env.EMAIL_PASS,
    },
});

async function sendConfirmation(toEmail, name, message) {
    try {
        await transporter.sendMail({
            from: `"LoginFrom"<${process.env.EMAIL_USER}>`,
            to: toEmail,
            subject: `Confirmation from LoginFrom`,
            html: `
            <h3>Hi ${name},</h3>
        <p>Thank you for submitting your message:</p>
        <blockquote>${message}</blockquote>
        <p>We'll get back to you soon!</p>
            `,
        });
        console.log(`Confirmation email sent to ${toEmail}`);
    } catch (err) {
        console.warn("Email sending failed:", err.message);
    }
}

module.exports = sendConfirmation;
