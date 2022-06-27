import { ChevronDownIcon } from '@chakra-ui/icons';
import { Box, Button, Container, Flex, Menu, MenuButton, MenuItem, MenuList } from '@chakra-ui/react';
import { PATH } from '@config/path';
import { APP } from '@utils/constants';
import { Link, useNavigate } from 'react-router-dom';

const Navbar = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    navigate(PATH.LOGIN);
  };

  return (
    <Box bg={'white'}>
      <Container maxWidth={'8xl'}>
        <Flex alignItems={'center'} justifyContent={'space-between'} boxShadow={{ base: 'sm' }} py={'4'}>
          <Box>
            <Button as={Link} to={PATH.DASHBOARD} variant={'none'} p={0} fontSize={{ base: 'xs', md: 'md' }}>
              {APP.name}
            </Button>
          </Box>
          <Flex gap={{ base: 2, md: 6 }}>
            <Button as={Link} to={PATH.DASHBOARD} variant={'none'} p={0} fontSize={{ base: 'xs', md: 'md' }}>
              Home
            </Button>
            <Menu>
              <MenuButton
                as={Button}
                variant={'none'}
                p={0}
                fontSize={{ base: 'xs', md: 'md' }}
                rightIcon={<ChevronDownIcon />}>
                John Doe
              </MenuButton>
              <MenuList fontSize={{ base: 'xs', md: 'md' }}>
                <MenuItem onClick={() => navigate(PATH.ACCOUNT)}>Akun</MenuItem>
                <MenuItem onClick={handleLogout}>Logout</MenuItem>
              </MenuList>
            </Menu>
          </Flex>
        </Flex>
      </Container>
    </Box>
  );
};

export default Navbar;
