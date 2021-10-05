import { ColorModeScript } from '@chakra-ui/react';
import React, { StrictMode } from 'react';
import ReactDOM from 'react-dom';
import AppRouter from './AppRouter';

ReactDOM.render(
  <StrictMode>
    <ColorModeScript />
    <AppRouter />
  </StrictMode>,
  document.getElementById('root')
);
