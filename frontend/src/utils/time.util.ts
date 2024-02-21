const formatTimeToLocaleString = (time: string | number) => {
  const date = new Date(time);

  console.log(date.toLocaleString());
  return date.toLocaleString();
};

export { formatTimeToLocaleString };
