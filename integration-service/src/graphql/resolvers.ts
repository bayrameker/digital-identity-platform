import { fetchThirdPartyData, postThirdPartyData } from '../services/apiService';

const resolvers = {
    Query: {
        getData: async (_: any, args: any) => {
            const data = await fetchThirdPartyData('https://api.thirdparty.com/data', args.params);
            return { success: true, data };
        },
    },
    Mutation: {
        postData: async (_: any, args: any) => {
            const data = await postThirdPartyData('https://api.thirdparty.com/data', args.input);
            return { success: true, data };
        },
    },
};

export default resolvers;
