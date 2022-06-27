import { ChevronDownIcon } from '@chakra-ui/icons';
import { Box, Button, Flex, Menu, MenuButton, MenuItem, MenuList } from '@chakra-ui/react';
import { PATH } from '@config/path';
import { APP } from '@utils/constants';
import { Link, useNavigate } from 'react-router-dom';

const Navbar = () => {
  let navigate = useNavigate();

  const handleLogout = () => {
    navigate(PATH.LOGIN);
  };
  return (
    <Flex
      bg={'white'}
      alignItems={'center'}
      justifyContent={'space-between'}
      boxShadow={{ base: 'sm' }}
      py={{ base: 2, md: 4 }}
      px={{ base: 2, md: 20 }}>
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
  );
};

export default Navbar;
