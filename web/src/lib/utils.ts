import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function removeMidAddress(address: string) {
  if (address.length <= 20) {
    return address;
  }
  return `${address.slice(0, 9)}...${address.slice(-10)}`;
}
