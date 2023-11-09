const convertLocationStrToNumber = (locationStr: string): number => {
  console.log("locationStr: ", locationStr);

  let loc = Number(locationStr);

  console.log("loc: ", loc);
  return loc;
};

export { convertLocationStrToNumber };
