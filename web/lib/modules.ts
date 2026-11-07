/** The numbered tiles on the home screen, from docs/screens.md. */
export type Module = {
  number: string;
  slug: string;
  title: string;
  icon: string;
};

export const MODULES: Module[] = [
  { number: "01", slug: "cash-finance", title: "Cash & Finance", icon: "cash" },
  { number: "02", slug: "equipment-assets", title: "Equipment & Asset Management", icon: "wrench" },
  { number: "03", slug: "general-controls", title: "General Controls", icon: "clipboard" },
  { number: "04", slug: "inventory-stores", title: "Inventory & Stores", icon: "box" },
  { number: "05", slug: "production-kitchen", title: "Production & Kitchen Operations", icon: "chef" },
  { number: "06", slug: "quality-compliance", title: "Quality, Food Safety & Compliance", icon: "shield" },
  { number: "07", slug: "sales-customer", title: "Sales & Customer Service", icon: "chat" },
  { number: "08", slug: "store-operations", title: "Store Operations", icon: "store" },
  { number: "09", slug: "human-resources", title: "Human Resources", icon: "people" },
  { number: "10", slug: "training", title: "Trainings & Product Knowledge", icon: "book" },
  { number: "11", slug: "security", title: "Security", icon: "lock" },
  { number: "12", slug: "sops", title: "Standard Operating Procedures (SOPs)", icon: "doc" },
  { number: "13", slug: "complexes", title: "Complexes", icon: "map" },
];
