import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Golidate",
  description: "Composable Go validation library",
  themeConfig: {
    search: {
      provider: "local",
    },
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "Home", link: "/" },
      { text: "Documentation", link: "/markdown-examples" },
      { text: "API Reference", link: "/markdown-examples" },
    ],

    sidebar: [
      {
        text: "Getting Started",
        items: [
          { text: "Installation", link: "/installation" },
          { text: "Defining Rules", link: "/installation" },
          { text: "Translating Results", link: "/installation" },
          { text: "Formatting Outputs", link: "/installation" },
        ],
      },
      {
        text: "Rules",
        items: [
          { text: "Validating Rules", link: "/available-rules" },
          { text: "Available Rules", link: "/available-rules" },
          { text: "Custom Rules", link: "/rules/explanation" },
          { text: "Rule Composition", link: "/composing-rules" },
        ],
      },
      {
        text: "Translations",
        items: [
          { text: "Translating Results", link: "/translation" },
          { text: "Available Languages", link: "/translation" },
        ],
      },
      {
        text: "Formatting",
        items: [
          { text: "Formatting Results", link: "/rules" },
          { text: "Available Formats", link: "/available-rules" },
          { text: "Custom Formats", link: "/rules/explanation" },
        ],
      },
    ],

    socialLinks: [
      { icon: "github", link: "https://github.com/vuejs/vitepress" },
    ],
  },
});
