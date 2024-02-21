/** @type {import('tailwindcss').Config} */
export const content = ["./view/**/*.html"];
export const theme = {
  extend: {
    colors: {
      dark: "#151515",
    },
    fontFamily: {
      manrope: ["Manrope", "sans-serif"],
    },
  },
};
export const plugins = [];
export const corePlugins = { preFlight: true };
