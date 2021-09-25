import { gql, useQuery, useMutation } from "urql";

export function Questions() {
  const [{ data, fetching, error }] = useQuery({ query });

  if (fetching) return "Loading...";
  if (error) return `Error: ${error}`;

  data.questions.sort(sortByWeight)
  data.questions.forEach(q => {
    if(q.options !== undefined) q.options.sort(sortByWeight)
  })

  return (
    <code>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </code>
  );
}

const sortByWeight = (a,b) => {
  return a.weight-b.weight
}

const query = gql`
  query {
    questions {
      __typename
      ... on ChoiceQuestion {
        id
        body
        weight
        options {
          id
          body
          weight
        }
      }
      ... on TextQuestion {
        id
        body
        weight
      }
    }
  }
`;

