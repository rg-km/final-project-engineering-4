import { ChevronDownIcon } from '@chakra-ui/icons';
import {
  Badge,
  Box,
  Button,
  Container,
  Flex,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Stack,
} from '@chakra-ui/react';
import { PATH } from '@config/path';
import { useAuthStore } from '@store/auth.store';
import { APP } from '@utils/constants';
import { Link, useNavigate } from 'react-router-dom';

const Navbar = () => {
  const navigate = useNavigate();

  const { role, user, handleLogout } = useAuthStore((state) => state);

  return (
    <Box bg={'white'}>
      <Container maxWidth={'8xl'}>
        <Flex alignItems={'center'} justifyContent={'space-between'} boxShadow={{ base: 'sm' }} py={'4'}>
          <Stack direction={'row'} alignItems={'center'} spacing={'3'}>
            <Button as={Link} to={PATH.DASHBOARD} variant={'none'} p={0} fontSize={{ base: 'xs', md: 'md' }}>
              {APP.name}
            </Button>
            <Badge>{role.title}</Badge>
          </Stack>
          <Flex gap={{ base: 2, md: 6 }}>
            <Button as={Link} to={PATH.DASHBOARD} variant={'none'} p={0} fontSize={{ base: 'xs', md: 'md' }}>
              Home
            </Button>
            <Menu>
              <MenuButton
                as={Button}
                variant={'none'}
                fontSize={{ base: 'xs', md: 'md' }}
                rightIcon={<ChevronDownIcon />}
                textTransform={'capitalize'}>
                {user.nama}
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
