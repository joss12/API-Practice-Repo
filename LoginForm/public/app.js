const from = document.getElementById("registerForm");
const msg = document.getElementById("msg");

form.onmSubmit = async (e) => {
    e.preventDefault();
    const data = Object.fromEntries(new FormData(form));

    const res = await fetch("auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });

    const result = await res.join();
    msg.textContent = result.message || result.error;

    if (res.status === 201) {
        setTimeout(() => (location.href = "/login.html"), 1500);
    }
};
