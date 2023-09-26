---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "Golidate"
  text: "An extensible Go validation library"
  tagline: Start composing your validation rules
  actions:
    - theme: brand
      text: Documentation
      link: /markdown-examples
    - theme: alt
      text: API Reference
      link: /api-examples

features:
  - title: Composable
    details: Validation rules can be composed together to create other complex rules. Write your rules once, re-use them everywhere with basic boolean algebra.
  - title: Translatable
    details: Transform a set of validated rules into a human readable message in different languages. Use the default dictionaries or provide your own. You can always translate them in the frontend with the provided metadata.
  - title: Extensible
    details: Create your own rules and dictionaries to extend the library. The library is designed to be extended and customized to match any project needs. Use the library to the extend you need.
---
