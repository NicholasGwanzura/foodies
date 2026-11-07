const paths: Record<string, React.ReactNode> = {
  home: <path d="M3 11.5 12 4l9 7.5M5.5 10v9.5h13V10M10 19.5v-6h4v6" />,
  tasks: <path d="M9 6h11M9 12h11M9 18h11M4 6l1 1 2-2M4 12l1 1 2-2M4 18l1 1 2-2" />,
  approve: <path d="M12 3l7.5 3v5.5c0 4.5-3 7.8-7.5 9.5-4.5-1.7-7.5-5-7.5-9.5V6zM8.5 12l2.5 2.5 4.5-5" />,
  grid: <path d="M4 4h7v7H4zM13 4h7v7h-7zM4 13h7v7H4zM13 13h7v7h-7z" />,
  bell: <path d="M18 9.5a6 6 0 1 0-12 0c0 5.5-2.5 7-2.5 7h17S18 15 18 9.5M10 20a2.2 2.2 0 0 0 4 0" />,
  user: <path d="M12 12a4.2 4.2 0 1 0 0-8.4 4.2 4.2 0 0 0 0 8.4M4.5 20.5c1-3.6 4-5.3 7.5-5.3s6.5 1.7 7.5 5.3" />,
  alert: <path d="M12 3.5 2.5 20h19zM12 10v4.5M12 17.3v.2" />,
  search: <path d="M10.8 18a7.2 7.2 0 1 0 0-14.4 7.2 7.2 0 0 0 0 14.4M16 16l4.5 4.5" />,
  cash: <path d="M3 7h18v10H3zM12 14.8a2.8 2.8 0 1 0 0-5.6 2.8 2.8 0 0 0 0 5.6M6 7v10M18 7v10" />,
  gear: <path d="M12 15a3 3 0 1 0 0-6 3 3 0 0 0 0 6M12 2.5v3M12 18.5v3M21.5 12h-3M5.5 12h-3M18.7 5.3l-2.1 2.1M7.4 16.6l-2.1 2.1M18.7 18.7l-2.1-2.1M7.4 7.4 5.3 5.3" />,
  clipboard: <path d="M9 4h6v3H9zM9 5.5H5.5v15h13v-15H15M9 12h6M9 16h4" />,
  box: <path d="M12 3 3.5 7.5v9L12 21l8.5-4.5v-9zM3.5 7.5 12 12l8.5-4.500M12 12v9" />,
  chef: <path d="M7 14.5c-2.5 0-4-1.8-4-3.800s1.7-3.8 4-3.700c.5-2 2.5-3.5 5-3.500s4.5 1.5 5 3.500c2.3-.1 4 1.7 4 3.700s-1.5 3.8-4 3.800V20H7zM7 16.800h10" />,
  shield: <path d="M12 3l7.5 3v5.5c0 4.5-3 7.8-7.5 9.5-4.5-1.7-7.5-5-7.5-9.5V6z" />,
  chat: <path d="M4 5h16v11H9l-5 4zM8 9.500h8M8 12.500h5" />,
  store: <path d="M4 10v10h16V10M3 10l2-6h14l2 6zM9.5 20v-6h5v6" />,
  people: <path d="M9 11.500a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7M2.5 20c.8-3 3.3-4.5 6.5-4.500s5.7 1.5 6.5 4.500M16 4.800a3.5 3.5 0 0 1 0 6.400M18 15.800c1.8.6 3 2 3.5 4.2" />,
  book: <path d="M12 6.500C10 5 7.5 4.5 4 4.500v14c3.5 0 6 .5 8 2 2-1.5 4.5-2 8-2v-14c-3.5 0-6 .5-8 2zM12 6.500v14" />,
  lock: <path d="M6 11h12v9.500H6zM8.5 11V7.500a3.5 3.5 0 0 1 7 0V11" />,
  doc: <path d="M6 3h8.500L19 7.500V21H6zM14 3v5h5M9 13h7M9 17h5" />,
  map: <path d="M12 21s7-5.5 7-11a7 7 0 1 0-14 0c0 5.5 7 11 7 11M12 12.500a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5" />,
  wrench: <path d="M20.5 6.500a5 5 0 0 1-6.8 5.800L7 19a2.1 2.1 0 0 1-3-3l6.7-6.700A5 5 0 0 1 16.5 2.500L13.5 5.500l2 3 3 1z" />,
  chevron: <path d="m9.5 5.5 6.5 6.5-6.5 6.5" />,
  plus: <path d="M12 5v14M5 12h14" />,
};

export function Icon({ name, size = 22 }: { name: string; size?: number }) {
  return (
    <svg
      width={size}
      height={size}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.9"
      strokeLinecap="round"
      strokeLinejoin="round"
      aria-hidden="true"
    >
      {paths[name] ?? paths.doc}
    </svg>
  );
}
