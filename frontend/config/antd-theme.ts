import type { ThemeConfig } from "antd";

export const antdTheme: ThemeConfig = {
  token: {
    colorPrimary: "#2563eb",
    colorSuccess: "#159a6a",
    colorWarning: "#d97706",
    colorError: "#dc2626",
    colorInfo: "#2563eb",
    colorText: "#172033",
    colorTextSecondary: "#536174",
    colorBgBase: "#eef4f8",
    colorBgContainer: "rgba(255, 255, 255, 0.86)",
    colorBorder: "rgba(113, 128, 150, 0.22)",
    borderRadius: 8,
    borderRadiusLG: 8,
    boxShadow: "0 16px 40px rgba(23, 32, 51, 0.08)",
    fontFamily:
      "Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif",
  },
  components: {
    Button: {
      borderRadius: 8,
      controlHeight: 38,
      fontWeight: 600,
    },
    Card: {
      borderRadiusLG: 8,
      paddingLG: 24,
    },
    Layout: {
      bodyBg: "transparent",
      headerBg: "transparent",
      siderBg: "transparent",
    },
    Menu: {
      itemBorderRadius: 8,
      itemHeight: 42,
      itemSelectedBg: "rgba(37, 99, 235, 0.12)",
      itemSelectedColor: "#1d4ed8",
    },
    Typography: {
      titleMarginBottom: 0,
      titleMarginTop: 0,
    },
  },
};
