const formatTimeToLocaleString = (time: string) => {
  const date = new Date(time);
  return date.toLocaleString();
};

export { formatTimeToLocaleString };
