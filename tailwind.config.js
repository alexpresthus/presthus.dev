/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["views/**/*.templ"],
    theme: {
        extend: {
            colors: {
                background: "#111111",
                primary: "#fae3d9",
                primaryGradient: "#f4c2ac",
                accent: "#61c0bf",
                accentGradient: "85cfce",
                text: {
                    primary: "#ffffff",
                },
            },
        },
    },
    plugins: [],
}


