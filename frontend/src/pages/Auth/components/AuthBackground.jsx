import { Box } from '@chakra-ui/react';
import React from 'react';

const AuthBackground = ({ children }) => {
  return (
    <Box
      bgImage={require('@images/picture/bg_auth.jpg')}
      height={'100vh'}
      display={'flex'}
      bgRepeat={'no-repeat'}
      bgPosition={'center'}
      bgSize={'cover'}
      objectFit={'cover'}>
      {children}
    </Box>
  );
};

export default AuthBackground;
