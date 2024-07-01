export const getNameInitials = (name: string, count = 2) => {
  const initials = name
    .split(" ")
    .map((n) => n[0])
    .join("");

  const filtered = initials.replace(/[^a-aZ-Z]/g, "");
  return filtered.slice(0, count).toUpperCase();
};
