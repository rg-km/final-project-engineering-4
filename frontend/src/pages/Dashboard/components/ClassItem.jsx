import { Link } from 'react-router-dom';

const { Box, Image, Heading, Code, Text } = require('@chakra-ui/react');

const ClassItem = ({ to, kelas }) => {
  return (
    <Box
      bg={'white'}
      p={{ base: 2, md: 4 }}
      as={Link}
      to={to}
      rounded={'md'}
      shadow={'sm'}
      border={'1px'}
      borderColor={'transparent'}
      transition={'all 200ms ease'}
      _hover={{ borderColor: 'primary.500' }}>
      <Box mb={{ base: 2, md: 4 }} bg={'blue.100'} p={{ base: 2 }} rounded={'md'}>
        <Image src={require('@images/icons/class.png')} alt={'plus icon'} m={'auto'} w={'60%'} />
      </Box>
      <Box>
        <Heading fontSize={{ base: 'xs', md: 'sm' }}>
          {kelas.mata_pelajaran} • {kelas.nama_kelas}
        </Heading>
        <Code fontSize={{ base: 'xs' }}>{kelas.kode_kelas}</Code>
        <Text fontSize={{ base: 'xs' }} mt={2}>
          {kelas.keterangan}
        </Text>
      </Box>
    </Box>
  );
};

export default ClassItem;
