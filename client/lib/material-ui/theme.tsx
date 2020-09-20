import { createMuiTheme } from '@material-ui/core/styles';

// Create a theme instance.
export const theme = createMuiTheme({
  palette: {
    background: {
      paper: 'rgb(27,43,55)',
      default: 'rgb(27,43,55)',
    },
    // Material UI のダークテーマを参考にした
    // https://material-ui.com/customization/default-theme/ {
    text: {
      primary: '#fff',
      secondary: 'rgba(255, 255, 255, 0.7)',
      disabled: 'rgba(255, 255, 255, 0.5)',
      hint: 'rgba(255, 255, 255, 0.5)',
    },
    divider: 'rgba(255, 255, 255, 0.12)',
    // }
  },
});
