// assets/js/navigation.js

const NAV_CONFIG = {
  topNav: [
    { label: "Getting Started", href: "../getting-started/" },
    { label: "Models", href: "../models/" },
    { label: "Fields", href: "../fields/" },
    { label: "Queries", href: "../queries/" },
    { label: "Schema", href: "../schema/" },
    { label: "Components", href: "../components/" },
    { label: "Best Practices", href: "../best-practices/" },
  ],

  sidebars: {
    "getting-started": {
      title: "Getting Started",
      items: [
        { label: "Installation", href: "#installation" },
        { label: "Importing Package", href: "#importing-package" },
      ]
    },
    "models": {
      title: "Models",
      items: [
        { label: "Defining a Model", href: "#defining-model" },
        { label: "Field Types", href: "#field-types" },
        { label: "Field Modifiers", href: "#field-modifiers" },
        { label: "Accessing Fields", href: "#accessing-fields" },
      ]
    },

    "fields": {
      title: "Fields",
      items: [
        { label: "Creating Fields", href: "#creating-fields" },
        { label: "Field Types", href: "#field-types" },
        { label: "Field Modifiers", href: "#field-modifiers" },
        { label: "Defaults", href: "#defaults" },
        { label: "Indexes & Keys", href: "#indexes" },
        { label: "Examples", href: "#examples" },
      ]
    },

    "queries": {
      title: "Queries",
      items: [
        { label: "Starting Queries", href: "#starting-queries" },
        { label: "SELECT", href: "#select" },
        { label: "INSERT", href: "#insert" },
        { label: "UPDATE", href: "#update" },
        { label: "DELETE", href: "#delete" },
        { label: "WHERE Conditions", href: "#where" },
        { label: "AND / OR", href: "#logical" },
        { label: "Pagination", href: "#pagination" },
      ]
    },

    "schema": {
      title: "Schema",
      items: [
        { label: "Schema Synchronization", href: "#schema-sync" },
        { label: "Auto Migration", href: "#auto-migration" },
        { label: "Migration Flags", href: "#migration-flags" },
        { label: "Environment Guidelines", href: "#environment-guidelines" },
      ]
    },

    "components": {
      title: "Components",
      items: [
        { label: "What Are Components", href: "#what-are-components" },
        { label: "Component Structure", href: "#component-structure" },
        { label: "Disk Representation", href: "#disk-representation" },
        { label: "Sync Behavior", href: "#sync-behavior" },
        { label: "Migration Flags", href: "#component-migration-flags" },
        { label: "Usage Guidelines", href: "#usage-guidelines" },
      ]
    },
    "best-practices": {
      title: "Best Practices",
      items: [
        { label: "Model Organization", href: "#model-organization" },
        { label: "Field Naming", href: "#field-naming" },
        { label: "Query Optimization", href: "#query-optimization" },
        { label: "Schema Management", href: "#schema-management" },
      ]
    }
  }
};

/* ------------------------
   RENDERERS
------------------------- */

function renderTopNav(activeSection) {
  const nav = document.getElementById("dynamic-top-nav");
  if (!nav) return;

  nav.innerHTML = NAV_CONFIG.topNav
    .map(item => {
      const active = item.label.toLowerCase() === activeSection ? "active" : "";
      return `<li class="nav-item">
        <a class="nav-link ${active}" href="${item.href}">${item.label}</a>
      </li>`;
    })
    .join("");
}

function renderSidebar(sectionKey) {
  const sidebar = document.getElementById("dynamic-sidebar");
  if (!sidebar || !NAV_CONFIG.sidebars[sectionKey]) return;

  const section = NAV_CONFIG.sidebars[sectionKey];

  sidebar.innerHTML = `
    <h6 class="sidebar-title">${section.title}</h6>
    <ul class="sidebar-nav">
      ${section.items.map(i =>
        `<li><a href="${i.href}">${i.label}</a></li>`
      ).join("")}
    </ul>
    <h6 class="sidebar-title mt-4">Other Sections</h6>
    <ul class="sidebar-nav">
      ${Object.keys(NAV_CONFIG.sidebars).map(key => {
        if (key === sectionKey) return "";
        const sec = NAV_CONFIG.sidebars[key];
        return `<li><a href="../${key}/">${sec.title}</a></li>`;
      }).join("")}
    </ul>
  `;
}

/* ------------------------
   INIT
------------------------- */

document.addEventListener("DOMContentLoaded", () => {
  const section = document.body.dataset.section;
  if (!section) return;

  renderTopNav(section);
  renderSidebar(section);
});
