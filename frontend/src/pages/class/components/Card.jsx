import { Box, Divider } from "@chakra-ui/react";

const Card = ({title, children}) => {
    return <Box
        bg={'white'}
        py={{ base: 1, md: 2 }}
        px={{ base: 2, md: 4 }}
        rounded={{ base: "md", md: "xl" }}
        shadow={'sm'}>
        <Box fontWeight={"bold"} fontSize={{ base: "xs", md: "md" }}>
            {title}
        </Box>
        <Divider my={{ base: 1 }} />
        <Box py={{base: 1, md: 4}}>
            {children}
        </Box>
    </Box>
}

export default Card;