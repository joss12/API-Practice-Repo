module.exports = function validateForm(name, email, message) {
    const errors = {};
    if (!name || name.length < 2) errors.name = "Name is too short";
    if (!email || !email.includes("@")) errors.email = "Invalid email";
    if (!message || message.length < 5) errors.message = "Message too short";

    return {
        valid: Object.keys(errors).length === 0,
        errors,
    };
};
