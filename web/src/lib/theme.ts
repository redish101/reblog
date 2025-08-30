import {
  BrandVariants,
  createDarkTheme,
  createLightTheme,
  Theme,
} from "@fluentui/react-components";

const variants: BrandVariants = {
  10: "#020404",
  20: "#111B1B",
  30: "#182E2D",
  40: "#1D3B3A",
  50: "#214947",
  60: "#255755",
  70: "#296564",
  80: "#2D7472",
  90: "#328381",
  100: "#4C908E",
  110: "#639D9B",
  120: "#79ABA9",
  130: "#8FB8B6",
  140: "#A5C6C4",
  150: "#BAD3D2",
  160: "#D0E1E0",
};

const baseThemeExtras: Partial<Theme> = {
  borderRadiusSmall: "8px",
  borderRadiusMedium: "8px",
  borderRadiusLarge: "8px",
  borderRadiusXLarge: "12px",
};

const lightTheme: Theme = {
  ...createLightTheme(variants),
  ...baseThemeExtras,
};

const darkTheme: Theme = {
  ...createDarkTheme(variants),
  ...baseThemeExtras,
};

darkTheme.colorBrandForeground1 = variants[110];
darkTheme.colorBrandForeground2 = variants[120];

export { lightTheme, darkTheme };
