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

    @Override
    public void start(Stage primaryStage) throws Exception {
        flowPane = new FlowPane();
        Button directButton = new Button("Schedule directly");
        Button centerButton = new Button("FxRunCenter");

        directButton.setOnMouseClicked(e -> directSchedule());
        centerButton.setOnMouseClicked(e -> centerSchedule());

        primaryStage.setScene(new Scene(new VBox(
                new HBox(
                        directButton,
                        centerButton,
                        new ProgressBar()
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
            int index = i;
//            Platform.runLater(() -> {
            ExceptionUtil.uncheck(() -> Thread.sleep(10));
            flowPane.getChildren().add(new Button("direct-" + index));
//            });
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
}
