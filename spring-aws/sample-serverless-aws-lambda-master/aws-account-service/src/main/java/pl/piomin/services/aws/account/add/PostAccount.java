package pl.piomin.services.aws.account.add;

import com.amazonaws.regions.Region;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDBClient;
import com.amazonaws.services.dynamodbv2.datamodeling.DynamoDBMapper;
import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.LambdaLogger;
import com.amazonaws.services.lambda.runtime.RequestHandler;

import pl.piomin.services.aws.account.model.Account;

public class PostAccount implements RequestHandler<Account, Account> {

	private DynamoDBMapper mapper;
	
	public PostAccount() {
		AmazonDynamoDBClient client = new AmazonDynamoDBClient();
		client.setRegion(Region.getRegion(Regions.US_EAST_1));
		mapper = new DynamoDBMapper(client);
	}
	
	@Override
	public Account handleRequest(Account a, Context ctx) {
		LambdaLogger logger = ctx.getLogger();
		mapper.save(a);		
		Account r = a;
		logger.log("Account: " + r.getId());
		return r;
	}

}
