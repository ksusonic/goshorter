document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("shorten-form");
    const resultDiv = document.getElementById("result");

    form.addEventListener("submit", async (event) => {
        event.preventDefault();

        const urlInput = document.getElementById("url");
        const url = urlInput.value;

        if (!url) {
            resultDiv.textContent = "Введите URL";
            return;
        }

        try {
            const response = await fetch("/api/shorten", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({url})
            });

            if (!response.ok) {
                message = (await response.json()).message || response.statusText;
                throw new Error(`${message}`);
            }

            const data = await response.json();
            resultDiv.innerHTML = `<a href="http://${data.shortened_url}" target="_blank">${data.shortened_url}</a>`;
        } catch (error) {
            resultDiv.textContent = `Ошибка: ${error.message}`;
        }
    });
});
