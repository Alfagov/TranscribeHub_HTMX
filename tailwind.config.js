/** @type {import('tailwindcss').Config} */
const plugin = require('tailwindcss/plugin')
module.exports = {
    darkMode: 'media',
    content: [
        "./templates/*.html",
        "./templates/**/*.html",
        "./node_modules/flowbite/**/*.js",
    ],
    theme: {
        extend: {},
    },
    daisyui: {
      themes: [
          "light",
          "dark",
          "night",
      ]
    },
    plugins: [
        require("daisyui"),
        plugin(function ({addVariant}) {
            addVariant('htmx-settling', ['&.htmx-settling', '.htmx-settling &'])
            addVariant('htmx-request', ['&.htmx-request', '.htmx-request &'])
            addVariant('htmx-swapping', ['&.htmx-swapping', '.htmx-swapping &'])
            addVariant('htmx-added', ['&.htmx-added', '.htmx-added &'])
        }),
        plugin(function ({addUtilities}) {
            addUtilities({
                /* Hide scrollbar for Chrome, Safari and Opera */
                '.no-scrollbar::-webkit-scrollbar': {
                    'display': 'none'
                },

                /* Hide scrollbar for IE, Edge and Firefox */
                '.no-scrollbar': {
                    '-ms-overflow-style': 'none', /* IE and Edge */
                    'scrollbar-width': 'none'  /* Firefox */
                }
            })
        })
    ],
}

