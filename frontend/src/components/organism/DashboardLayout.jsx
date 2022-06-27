import { Box, Container } from '@chakra-ui/react';
import { Outlet } from 'react-router';
import Navbar from './Navbar';

const Dashboard = () => {
  return (
    <Box bg={'gray.100'} minH={'100vh'}>
      <Navbar />

      <Container maxWidth={'8xl'}>
        <Outlet />
      </Container>
    </Box>
  );
};

export default Dashboard;
