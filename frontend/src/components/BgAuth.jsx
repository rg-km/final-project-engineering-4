import { Box } from '@chakra-ui/react';
import React from 'react';
import bg_auth from "@images/bg_auth.jpg";
import PropTypes from 'prop-types';

const BgAuth = ({ children }) => {
    return <Box bg={'gray.50'} h={'100vh'} display={'flex'} bgImg={bg_auth} bgRepeat={'no-repeat'} bgSize={'cover'} objectFit={'cover'}>
        {children}
    </Box>
}

BgAuth.propTypes = {
    children: PropTypes.node
}

export default BgAuth;
