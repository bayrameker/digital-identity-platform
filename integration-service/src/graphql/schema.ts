import { gql } from 'apollo-server-express';

const typeDefs = gql`
  type Query {
    getData(params: DataParams): DataResponse
  }

  type Mutation {
    postData(input: DataInput): DataResponse
  }

  input DataParams {
    key: String
    value: String
  }

  input DataInput {
    key: String!
    value: String!
  }

  type DataResponse {
    success: Boolean!
    data: String
  }
`;

export default typeDefs;
