package com.github.acexy.tech.crypto.spacemesh.test;

import com.github.acexy.tech.crypto.spacemesh.proto.TransactionServiceGrpc;
import com.github.acexy.tech.crypto.spacemesh.proto.TxTypes;
import com.github.acexy.tech.crypto.spacemesh.proto.Types;
import com.google.protobuf.ByteString;
import io.grpc.Grpc;
import io.grpc.InsecureChannelCredentials;
import io.grpc.ManagedChannel;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

import java.util.Scanner;

public class ClientTest {


    private static String srv;

    @BeforeAll
    public static void beforeAll() {
        Scanner scanner = new Scanner(System.in);
        System.out.println("set your spacemesh node server: ");
        srv = scanner.nextLine();
    }

    private TransactionServiceGrpc.TransactionServiceBlockingStub transactionService() {
        ManagedChannel channel = Grpc.newChannelBuilder(srv, InsecureChannelCredentials.create()).build();
        return TransactionServiceGrpc.newBlockingStub(channel);
    }

    public static byte[] hexToByteArray(String inHex) {
        int len = inHex.length() / 2;
        byte[] result = new byte[len];
        for (int i = 0; i < len; i++) {
            result[i] = (byte) Integer.parseInt(inHex.substring(i * 2, i * 2 + 2), 16);
        }
        return result;
    }

    @Test
    public void transactionsState() {
        Types.TransactionId t = Types.TransactionId.newBuilder().setId(ByteString.copyFrom(hexToByteArray("ad0fa1fec81bbfa0700bdc945c5b8549a375c0b6d825b107dee3c5a53521abe6"))).build();
        TxTypes.TransactionsStateResponse resp = transactionService().transactionsState(TxTypes.TransactionsStateRequest.newBuilder().addTransactionId(t).setIncludeTransactions(true).build());
        System.out.println(resp.toString());
    }
}
