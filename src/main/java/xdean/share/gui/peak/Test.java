package xdean.share.gui.peak;

import javafx.application.Application;
import javafx.application.Platform;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.ProgressBar;
import javafx.scene.control.ScrollPane;
import javafx.scene.layout.FlowPane;
import javafx.scene.layout.HBox;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;
import xdean.jex.util.lang.ExceptionUtil;

public class Test extends Application {
    public static void main(String[] args) {
        launch(args);
    }

    FlowPane flowPane;
    Button stubButton;

    @Override
    public void start(Stage primaryStage) throws Exception {
        flowPane = new FlowPane();
        stubButton = new Button("stub");
        Button directButton = new Button("Large Work");
        Button centerButton = new Button("FxRunCenter Peak Clipping");
        Button directSetButton = new Button("Duplicate Work");
        Button centerSetButton = new Button("FxRunCenter De-duplication");

        directButton.setOnMouseClicked(e -> directSchedule());
        centerButton.setOnMouseClicked(e -> centerSchedule());
        directSetButton.setOnMouseClicked(e -> directSetText());
        centerSetButton.setOnMouseClicked(e -> centerSetText());

        primaryStage.setScene(new Scene(new VBox(
                new HBox(
                        directButton,
                        centerButton
                ),
                new HBox(
                        directSetButton,
                        centerSetButton
                ),
                new HBox(
                        new ProgressBar(),
                        stubButton
                ),
                new ScrollPane(flowPane)
        )));
        primaryStage.setWidth(500);
        primaryStage.setHeight(500);
        primaryStage.show();
    }

    private void directSchedule() {
        flowPane.getChildren().clear();
        for (int i = 0; i < 500; i++) {
            ExceptionUtil.uncheck(() -> Thread.sleep(10));
            flowPane.getChildren().add(new Button("direct-" + i));
        }
    }

    private void centerSchedule() {
        flowPane.getChildren().clear();
        for (int i = 0; i < 500; i++) {
            int index = i;
            FxRunCenter.runLater(() -> {
                ExceptionUtil.uncheck(() -> Thread.sleep(10));
                flowPane.getChildren().add(new Button("center-" + index));
            });
        }
    }

    private void directSetText() {
        for (int i = 0; i < 500; i++) {
            ExceptionUtil.uncheck(() -> Thread.sleep(10));
            stubButton.setText("direct-" + Math.random());
        }
    }

    private void centerSetText() {
        for (int i = 0; i < 500; i++) {
            int index = i;
            FxRunCenter.builder().id(stubButton).run(() -> {
                ExceptionUtil.uncheck(() -> Thread.sleep(10));
                stubButton.setText("center-" + Math.random());
            });
        }
    }
}
