import { extendTheme, withDefaultColorScheme } from '@chakra-ui/react';

const theme = extendTheme(
  {
    colors: {
      primary: {
        100: '#fce2db',
        200: '#f8c4b8',
        300: '#f5a794',
        400: '#f18971',
        500: '#ee6c4d',
        600: '#be563e',
        700: '#8f412e',
        800: '#5f2b1f',
        900: '#30160f',
      },
      secondary: {
        100: '#d4d6d9',
        200: '#a9adb3',
        300: '#7f848d',
        400: '#545b67',
        500: '#293241',
        600: '#212834',
        700: '#191e27',
        800: '#10141a',
        900: '#080a0d',
      },
    },
    fonts: {
      heading: '"Montserrat", sans-serif',
      body: '"Plus Jakarta Sans", sans-serif',
      mono: '"Fira Code", monospace',
    },
  },
  withDefaultColorScheme({ colorScheme: 'primary' })
);

export default theme;
